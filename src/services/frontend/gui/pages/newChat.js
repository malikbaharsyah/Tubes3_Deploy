import React, { useState } from 'react';
import styles from '../styles/NewChat.module.css';

export const NewChat = () => {
  const [clicked, setClicked] = useState(false);

  const handleClick = () => {
    setClicked(true);
    // Tambahkan fungsi yang ingin dijalankan saat button diklik di sini
  }

  return (
    <div className={styles['new-chat']} onClick={handleClick} >
      {'+ New Chat'}
    </div>
  )
}
