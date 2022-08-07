import React, {useEffect, useState} from 'react';
import axios from 'axios';
import logo from './logo.svg';
import useFetch from "react-fetch-hook"

import './App.css';

function App() {

  const [getRes, setGetResponse] = useState([])

  useEffect(()=>{
    axios.get("http://localhost:10000/")
    .then((res)=>{
      setGetResponse(res)
    })
  }, [])

  return(
  
  <div>
    <p>{getRes.data}</p>

  </div>
    
  )
}

export default App;

