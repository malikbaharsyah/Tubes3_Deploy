import React, { useState } from 'react';
import styles from '../styles/ChatBox.module.css';

export const TextBox = () => {
  const [text, setText] = useState('');
  const [messages, setMessages] = useState([]);
  const [lastSender, setLastSender] = useState('');

  const handleTextChange = (event) => {
    setText(event.target.value);
  };

  const handleSubmit = (event) => {
    event.preventDefault();
    if (text.trim() !== '') {
      setMessages([...messages, { text, sender: 'me' }]);
      setText('');
      setLastSender('me');
    }
  };

  const handleKeyDown = (event) => {
    if (event.key === 'Enter') {
      handleSubmit(event);
    }
  };

  return (
    // <div className={styles['chat-box']} style={{ overflowY: 'scroll' }}>
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
          {message.sender === 'me' && (
            <div className={`${styles.answerBubble} ${styles.bot}`}>
              ada yang bisa saya bantu?
            </div>
          )}
        </div>
      ))}
    </div>
    // </div>
  );
};

