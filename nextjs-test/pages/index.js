import Head from 'next/head'
import styles from '../styles/Home.module.css'

import TopBar from './topBar';
import Posts from './posts';

export default function Home() {
    return (
        <div>
            <TopBar/>
            <Posts/>
        </div>

    )
}
