import React from 'react';
import { Slider } from 'primereact/slider';
import { FetchData } from './MainMap'

function GSlider(setCircList) {
    const [value, setValue] = React.useState(0);
    const step = 12.5;
    let curTime;

    const calculateTime = () => {
        const hours = 24*parseFloat(value)/100;
        const minutes = Math.round(0);
        const seconds = Math.round(0);
        curTime = `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(seconds).padStart(2, '0')}`;
        return curTime;
    };

    return (
        <div className='gslid'>
            <Slider value={value} onChange={(e) => {
                setValue(e.value);
            }} step={step} />
            <p className="time">{calculateTime()}</p>
        </div>
    );
}

export default GSlider;
