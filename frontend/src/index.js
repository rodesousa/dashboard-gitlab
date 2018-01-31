import React from 'react';
import ReactDOM from 'react-dom';
import './index.css';
import Accordeon from './Accordeon';
import registerServiceWorker from './registerServiceWorker';

ReactDOM.render(<Accordeon />, document.getElementById('root'));
registerServiceWorker();
