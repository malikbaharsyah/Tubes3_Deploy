import React, { useState } from 'react';
import styles from '../styles/HistoryBox.module.css';

export const HistoryBox = () => {
    const buttons = ["We Go Shout Shout", "When we together", "You make me better", "My life without you", "Is A Missery"];
    const [clicked, setClicked] = useState(Array(buttons.length).fill(false));
  
    const handleClick = (index) => {
      const newClicked = [...clicked];
      newClicked[index] = true;
      setClicked(newClicked);
      // Tambahkan fungsi yang ingin dijalankan saat button diklik di sini
    };
  
    for (let i = 0; i < buttons.length; i++) {
      buttons[i] = (
        <div className={styles['historyBox']} onClick={() => handleClick(i)} style={{ top: `${125 + (i + 1) * 50}px` }}>
          {buttons[i]} 
        </div>
      );
    }
  
    return <div className={styles['historyBoxContainer']}>{buttons}</div>;
  };
  