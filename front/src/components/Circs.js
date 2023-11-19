import {Marker, Popup, Circle} from 'react-leaflet';
import L from 'leaflet';

var Circ = (props) => {
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
    }
    const pinMB = L.icon({
        // iconUrl: '',
        iconUrl: 'https://unpkg.com/leaflet@1.7.1/dist/images/marker-icon.png',
        iconSize: [24, 41],
        iconAnchor: [0, 44],
        popupAnchor: [12, -40],
        shadowUrl: null,
        shadowSize: null,
        shadowAnchor: null
    });
    return (
        <div>
            <Circle
                center={[props.x,props.y]}
                radius={props.r}
                stroke={false}
                fillOpacity={0.1}
                color={colors[props.name]}
            />
        </div>
    );
}

export default Circ;