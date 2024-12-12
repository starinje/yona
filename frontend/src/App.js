import React from 'react';
import UserList from './components/UserList';
import './App.css';

function App() {
  return (
    <div className="App">
      <div className="container">
        <h1>Yona User Management</h1>
        <UserList />
      </div>
    </div>
  );
}

export default App;
