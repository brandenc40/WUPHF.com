import React from 'react';
import './App.css';
import { WuphfForm } from './components';
import { formOptions } from './constants/formOptions';

function App() {
  return (
    <div className='App'>
      <WuphfForm formOptions={formOptions}></WuphfForm>
    </div>
  );
}

export default App;
