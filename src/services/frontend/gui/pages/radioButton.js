import React, { useState, useEffect } from 'react';
import styles from '../styles/RadioButton.module.css';
import HistoryBox from './history';

export const RadioButton = () => {
  const [selectedOption, setSelectedOption] = useState('');

  const handleOptionChange = (event) => {
    setSelectedOption(event.target.value);
    sessionStorage.setItem('selectedOption', event.target.value);
  };

  return (
    <div>
      <label className={styles.container}>
        <input type="radio" name="option" value="option1" onChange={handleOptionChange} />
        <span className={styles.checkmark}></span>
        Knuth Morris Pratt
      </label>
      <label className={styles.container}>
        <input type="radio" name="option" value="option2" onChange={handleOptionChange} />
        <span className={styles.checkmark}></span>
        Boyer Moore
      </label>
    </div>
  );
};

export default RadioButton;
