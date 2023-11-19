import { Marker } from 'react-leaflet';
import L from 'leaflet';
import Circ from '../components/Circs';

function loadJsonData(url, time) {
  console.log("huidate", `02.01.2006 ${time}`)
  return fetch(url, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      date: `15.08.2021 ${time}`,
    }),
  })
    .then(response => response.json())
    .catch(error => console.error('Error loading JSON data:', error));
}

function createComponent(props) {
  const pinMB = L.icon({
    iconUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png',
    iconSize: [24, 41],
    iconAnchor: [0, 44],
    popupAnchor: [12, -40],
    shadowUrl: null,
    shadowSize: null,
    shadowAnchor: null
  });
  return (
    <div className='i-want-to-fucking-kill-myself'>
      <Circ x={props.x} y={props.y} r={100} name={props.name} />
    </div>
  );
}

export { loadJsonData, createComponent };
