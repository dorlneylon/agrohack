import React, { useEffect, useState } from 'react';
import { MapContainer, TileLayer } from 'react-leaflet';
import GSlider from './Slider';
import Circ from './Circs';
import Picker from './Picker';
import { loadJsonData, createComponent } from '../service/Loader';

let circList, setCircList;
const fetchData = async (setCircList, time) => {
  try {
    const data = await loadJsonData('http://localhost:8081/points/getAll', time);
    // Assuming data.Data is an array of points
    console.log(data.data.length);
    const newCircList = data.data.map((point, index) => (
        // <div key={index}>{point.name}</div>
        // createComponent(point)
        // <Circ x={point.x} y={point.y} r="1000" name={point.name} />

        <Circ key = {index} x={point.x} y={point.y} r="7000" name={point.name} />
    ));
    setCircList(newCircList);
  } catch (error) {
    console.error('Error fetching data:', error);
  }
};

const FetchData = async (time) => {
  fetchData(setCircList, time)
}



const MainMap = () => {
  const center = [46.724727, 38.277663];
  const zoom = 13;

  [circList, setCircList] = useState([]);

  useEffect(() => {
  }, []); // The empty dependency array ensures that this effect runs only once after the initial render.

  return (
      <div className="lilroot">
        {/* <GSlider /> */}
        <Picker />
        {/* <div className='sec'>
        <Picker />
        </div> */}
        <MapContainer center={center} zoom={zoom} scrollWheelZoom={false} style={{ height: '100vh', width: '100wh' }}>
          <TileLayer
              attribution='&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a> contributors &copy; <a href="https://carto.com/attributions">CARTO</a>'
              url='https://{s}.basemaps.cartocdn.com/dark_all/{z}/{x}/{y}{r}.png'
          />
          {/* Render Circ components from the state */}
          {circList}
        <Circ x={center[0]} y={center[1]} r = "1000"/>
        </MapContainer>
      </div>
  );
};

export default MainMap;
export {FetchData};

