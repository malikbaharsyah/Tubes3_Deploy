import React, { useState, useEffect } from 'react';
import styles from '../styles/HistoryBox.module.css';
import axios from 'axios';

export const HistoryBox = ({ onClearClick, setMessagesToList,
  isSubmitClicked, setIsSubmitClicked, userInput,
  addMessage}) => {
  const buttons = ["History 1", "History 2", "History 3", "History 4", "History 5"];
  const [activeIndex, setActiveIndex] = useState(-1);

  useEffect(() => {
    const performSubmitAction = async () => {
      if (isSubmitClicked) {
        // Clear the history or perform any other desired actions
        if (activeIndex === -1) {
          alert("Pilih history terlebih dahulu!");
        } else {
          const botResponse = await addMessage(userInput);
          if (botResponse) {
          addMessageToHistory(activeIndex, userInput, botResponse);
          }
        }
        setIsSubmitClicked(false); // Reset the submit click state
      }
    };
  
    performSubmitAction();
  }, [isSubmitClicked, onClearClick]);

  const handleHistoryClick = (index) => {
    setActiveIndex(index);
    getHistoryDatabase(index);
    // Call the onHistoryClick callback function with the question and answer
    // setMessagesToList(listOfQuestions, listOfAnswers);
  };

  const handleClearClick = (index) => {
    setActiveIndex(-1);
    deleteHistoryDatabase(index);
    onClearClick();
  };

  const getHistoryDatabase = async (index) => {
    try {
      const response = await axios.get(`http://localhost:8000/api/history/${index+1}`);
      const database = response.data.history;
      if (database && database.length > 0) {
        var listOfQuestions = [];
        var listOfAnswers = [];
        for (var i = 0; i < database.length; i++) {
          listOfQuestions.push(database[i].pertanyaan);
          listOfAnswers.push(database[i].jawaban);
        }
        setMessagesToList(listOfQuestions, listOfAnswers);
      } else {
        // If the database is empty, set the history to empty arrays
        setMessagesToList([], []);
      }
    } catch (error) {
      console.error(error);
    }
  };
  

  const deleteHistoryDatabase = async (index) => {
    try {
      const response = await axios.delete(`http://localhost:8000/api/history/${index+1}`);
    } catch (error) {
      console.error(error);
    }
  }

  const addMessageToHistory = async (index, question, answer) => {
    try {
      const response = await axios.post(`http://localhost:8000/api/history/${index+1}`, {
        pertanyaan: question,
        jawaban: answer
      });
    } catch (error) {
      console.error(error);
    }
  }

  const historyButtons = buttons.map((button, index) => (
    <div
      className={`${styles['historyBox']} ${activeIndex === index ? styles['active'] : ''}`}
      style={{ top: `${155 + (index + 1) * 50}px` }}
      key={index}
    >
      <div className={styles['historyButton']} onClick={() => handleHistoryClick(index)}>
        {button}
      </div>
      {activeIndex === index && (
        <button className={styles['clearButton']} onClick={() => handleClearClick(index)}>
          Clear
        </button>
      )}
    </div>
  ));

  return <div className={styles['historyBoxContainer']}>{historyButtons}</div>;
};
