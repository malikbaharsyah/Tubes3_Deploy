import React, { useState, useEffect } from 'react';
import styles from '../styles/HistoryBox.module.css';
import axios from 'axios';

export const HistoryBox = ({ onClearClick, setMessagesToList,
  isSubmitClicked, setIsSubmitClicked, userInput,
  addMessage }) => {
  const buttons = ["History 1", "History 2", "History 3", "History 4", "History 5"];
  const [activeIndex, setActiveIndex] = useState(-1);
  const listOfQuestions = ["Q1", "Q2", "Q3", "Q4", "Q5"];
  const listOfAnswers = ["A1", "A2", "A3", "A4", "A5"];

  useEffect(() => {
    // Perform actions when the submit button is clicked
    if (isSubmitClicked) {
      // Clear the history or perform any other desired actions
      if (activeIndex === -1) {
        alert("Pilih history terlebih dahulu!")
      }
      else {
        addMessage(userInput);
      }
      setIsSubmitClicked(false); // Reset the submit click state
    }
  }, [isSubmitClicked, onClearClick]);

  const handleHistoryClick = (index) => {
    setActiveIndex(index);
    getHistoryDatabase(index);
    // Call the onHistoryClick callback function with the question and answer
    // setMessagesToList(listOfQuestions, listOfAnswers);
  };

  const handleClearClick = (index) => {
    setActiveIndex(-1);
    onClearClick();
  };

  const getHistoryDatabase = async (index) => {
    try {
      const response = await axios.get(`http://localhost:8000/api/history/${index+1}`);
      const database = response.data.history;
      if (database) {
        var listOfQuestions = [];
        var listOfAnswers = [];
        for (var i = 0; i < database.length; i++) {
          listOfQuestions.push(database[i].pertanyaan);
          listOfAnswers.push(database[i].jawaban);
        }
        setMessagesToList(listOfQuestions, listOfAnswers);
      }
    } catch (error) {
      console.error(error);
    }
  }

  const historyButtons = buttons.map((button, index) => (
    <div
      className={`${styles['historyBox']} ${activeIndex === index ? styles['active'] : ''}`}
      style={{ top: `${125 + (index + 1) * 50}px` }}
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
