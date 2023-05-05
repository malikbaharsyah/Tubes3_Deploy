import React, { useState } from 'react';
import styles from '../styles/RadioButton.module.css';
import { TextBox } from './chatBox';

export const RadioButton = () => {
  const [selectedOption, setSelectedOption] = useState('');

  const handleOptionChange = (event) => {
    const value = event.target.value;
    if (value === selectedOption) {
      setSelectedOption('');
    } else {
      setSelectedOption(value);
    }
  };

  function getSelectedOption() {
    return selectedOption;
  }

  return (
    <div>
      <label className={styles.container}>
        <input type="radio" name="option" value="option1" checked={selectedOption === 'option1'} onChange={handleOptionChange} />
        <span className={styles.checkmark}></span>
        Knuth Morris Pratt
      </label>
      <label className={styles.container}>
        <input type="radio" name="option" value="option2" checked={selectedOption === 'option2'} onChange={handleOptionChange} />
        <span className={styles.checkmark}></span>
        Boyer Moore
      </label>
      <TextBox
        getSelectedOption={getSelectedOption}
      />
    </div>
  );
};
