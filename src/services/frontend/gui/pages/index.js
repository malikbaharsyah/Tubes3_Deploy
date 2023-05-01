import Head from 'next/head';
import styles from '../styles/Home.module.css';
import { LeftSideBar } from '../styles/Home.module.css';

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
        <div className={styles['home-atas']}>
        </div>
      </div>
    </div>
  );
}