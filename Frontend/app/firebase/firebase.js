import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";

const firebaseConfig = {
    apiKey: "AIzaSyBScjFY9BNrKnjkqGzp2ZBr_Mi0nB4dIPE",
    authDomain: "clickcraft-ac100.firebaseapp.com",
    projectId: "clickcraft-ac100",
    storageBucket: "clickcraft-ac100.appspot.com",
    messagingSenderId: "442600681149",
    appId: "1:442600681149:web:ef8549e048b0701d5251d4"
};


const app = initializeApp(firebaseConfig);

export const auth = getAuth(app);
export default app;