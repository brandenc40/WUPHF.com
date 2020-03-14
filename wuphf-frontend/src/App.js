import React from 'react';
import './App.css';
import { WuphfForm, PaypalDonate } from './components';
import { formOptions } from './constants/formOptions';

function App() {
  return (
    <div className='App'>
      <WuphfForm formOptions={formOptions}></WuphfForm>
      <div className='donate'>
        <PaypalDonate></PaypalDonate>
      </div>
    </div>
  );
}

export default App;
