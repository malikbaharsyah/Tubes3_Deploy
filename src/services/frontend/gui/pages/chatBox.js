import React, { useState, useEffect } from 'react';
import styles from '../styles/ChatBox.module.css';
import axios from 'axios';
import { HistoryBox } from './history';

export const TextBox = ({getSelectedOption}) => {
  const [text, setText] = useState('');
  const [messages, setMessages] = useState([]);
  const [lastSender, setLastSender] = useState('');
  const [botmessages, setBotMessages] = useState([]);
  const [newBotMessage, setNewBotMessage] = useState({});
  const [isSubmitClicked, setIsSubmitClicked] = useState(false);

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  const handleSubmit = async (event) => {
    event.preventDefault();
    setIsSubmitClicked(true);
  };
  

  const addMessage = async (text) => {
    if (text.trim() !== '') {
      const newMessage = { text, sender: 'me' };
      setMessages([...messages, newMessage]);
      setText('');
      setLastSender('me');
      console.log(getSelectedOption);
      try {
        var response;
        if (getSelectedOption === "option1") {
          response = await axios.get(`http://localhost:8000/api/gpt/0/${text}`);
        } else {
          response = await axios.get(`http://localhost:8000/api/gpt/1/${text}`);
        }
        const database = response.data.answer[0];
        if (database) {
          const newBotMessage = { text: database, sender: 'bot' };
          setBotMessages([...botmessages, newBotMessage]);
          setLastSender('bot');
          return newBotMessage.text;
        }
      } catch (error) {
        console.error(error);
      }
      const newBotMessage = { text: 'Maaf, saya tidak mengerti.', sender: 'bot' };
      setBotMessages([...botmessages, newBotMessage]);
      setLastSender('bot');
      return newBotMessage.text;
    }
  };
  

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleSubmit(event);
    }
  };


  const clearMessages = () => {
    setMessages([]);
    setBotMessages([]);
  };

  const updateMessages = (question, answer) => {
    setMessages(prevMessages => [...prevMessages, { text: question, sender: 'me' }]);
    setBotMessages(prevBotMessages => [...prevBotMessages, { text: answer, sender: 'bot' }]);
    setLastSender('bot');
  };  

  const setMessagesToList = (listOfQuestions, listOfAnswers) => {
    clearMessages();
    for (let i = 0; i < listOfQuestions.length; i++) {
      updateMessages(listOfQuestions[i], listOfAnswers[i]);
    }
  };


  return (
    <div className={styles.TextBox} style={{ overflowY: 'scroll' }}>
      <form onSubmit={handleSubmit} onKeyDown={handleKeyDown}>
        <input
          type="text"
          value={text}
          onChange={handleTextChange}
          placeholder="Send a message."
        />
        <div className={styles['button-send']}>
          <button type="submit">Send</button>
        </div>
      </form>
      {messages.map((message, index) => (
        <div key={index}>
          <div className={`${styles.bubble} ${styles[message.sender]}`}>
            {message.text}
            {message.sender === 'me' && <div className={styles.arrow} />}
          </div>
          {message.sender === 'me' && lastSender === 'bot' && (
            <div className={`${styles.answerBubble} ${styles.bot}`}>
              {botmessages[index].text}
            </div>
          )}
        </div>
      ))}
      <HistoryBox
      onClearClick={clearMessages}
      setMessagesToList={setMessagesToList}
      isSubmitClicked={isSubmitClicked}
      setIsSubmitClicked={setIsSubmitClicked}
      userInput={text}
      addMessage={addMessage}/>
    </div>
  );
};
