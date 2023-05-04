import Head from 'next/head';
import styles from '../styles/Home.module.css';
import leftSideBarStyles from '../styles/LeftSideBar.module.css';
import ChatBoxStyles from '../styles/ChatBox.module.css';
import RadioButtonStyles from '../styles/RadioButton.module.css';

import { TextBox } from './textBox';
import { RadioButton } from './radioButton';
import { NewChat } from './newChat';
 

export default function Home() {
  return (
    <div>
      <Head>
        <title>Halaman Utama</title>
        <link rel="icon" href="/favicon.ico" />
        <style>{`
          body {
            margin: 0;
            padding: 0;
            overflow: hidden; /* prevent scrolling */
          }
        `}</style>
      </Head>
      <div className={`${styles.column} ${styles['home']}`}>
        <RadioButton />
        <div className={`${leftSideBarStyles['left-sidebar']}`}>
          <div className={`${leftSideBarStyles['sidebar-text']}`}>
            <h2>HISTORY</h2>
          </div>
            <NewChat />
        </div>
        <div className={`${leftSideBarStyles['bottom-left-sidebar']}`}>
        </div>
          <div className={`${ChatBoxStyles['chat-box']}`}>
            <TextBox />
          </div>
      </div>


      <script>
        {`
          // Your JavaScript code here
          `}
      </script>
    </div>
  );
}
