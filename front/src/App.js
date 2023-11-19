import MainMap from './components/MainMap';
import 'leaflet/dist/leaflet.css';
import $ from 'jquery';
import {FetchData} from "./components/MainMap";
import { useState } from 'react';
import "primereact/resources/themes/lara-light-indigo/theme.css";
import './App.css';
import { Dropdown } from 'primereact/dropdown';

import React, { useEffect } from 'react';

function App() {
    let prevTime = "";
    $(document).ready(() => {
        setInterval(() => {
            console.log($("#\\:r1\\:").val());
            let parsedDate = new Date($("#\\:r1\\:").val());
            console.log(parsedDate);
            if (parsedDate) {
                let formattedDate = `${(parsedDate.getDate() + '').padStart(2, '0')}_${((parsedDate.getMonth() + 1) + '').padStart(2, '0')}_${parsedDate.getFullYear()}_${(parsedDate.getHours() + '').padStart(2, '0')}_${(parsedDate.getMinutes() + '').padStart(2, '0')}`;
                console.log(formattedDate);
            }

// Format the date as dd_mm_yyyy_hh_mm

            // let curTime = ""
            // console.log(curTime)
            // if(prevTime !== curTime) {
            //     prevTime = curTime;
            //     FetchData(curTime);
            // }

        }, 5000);}
    );

  var [selectedColor, setSelectedColor] = useState(null);
    const colors1 = [
        { name: 'Милдью или ложная мучнистая роса', code: 'NY' },
        { name: 'Rome', code: 'RM' },
        { name: 'London', code: 'LDN' },
        { name: 'Istanbul', code: 'IST' },
        { name: 'Paris', code: 'PRS' }
    ];

  const colors2 = [
    <div><div className='legenda-red'></div><p className='legenda-text'>Милдью или ложная мучнистая роса</p></div>,
    <div><div className='legenda-blue'></div><p className='legenda-text'>Оидиум</p></div>,
    <div><div className='legenda-brown'></div><p className='legenda-text'>Антракноз</p></div>,
    <div><div className='legenda-green'></div><p className='legenda-text'>Серая гниль</p></div>,
    <div><div className='legenda-orange'></div><p className='legenda-text'>Чёрная пятнистость</p></div>,
    <div><div className='legenda-crimson'></div><p className='legenda-text'>Чёрная гниль</p></div>,
    <div><div className='legenda-white'></div><p className='legenda-text'>Белая гниль</p></div>,
    <div><div className='legenda-purple'></div><p className='legenda-text'>Вертициллезное увядание</p></div>,
    <div><div className='legenda-yellow'></div><p className='legenda-text'>Альтернариоз</p></div>,
    <div><div className='legenda-cyan'></div><p className='legenda-text'>Фузариоз</p></div>,
    <div><div className='legenda-mangenta'></div><p className='legenda-text'>Краснуха</p></div>,
    <div><div className='legenda-gray'></div><p className='legenda-text'>Бактериальный рак</p></div>
  ];

  const colors = {
    "Милдью или ложная мучнистая роса": "red",
    "Оидиум" : "blue",
    "Антракноз" : "brown",
    "Серая гниль" : "green",
    "Чёрная пятнистость" : "orange",
    "Чёрная гниль" : "crimson",
    "Белая гниль" : "white",
    "Вертициллезное увядание" : "purple",
    "Альтернариоз" : "yellow",
    "Фузариоз" : "cyan",
    "Краснуха" : "mangenta",
    "Бактериальный рак" : "gray"
  };

  return (
    <div className="App">
      <Dropdown className='legenda' value={selectedColor} options={colors2} placeholder='Легенда' />
      <MainMap />
    </div>
  );
}

export default App;

