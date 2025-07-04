import { BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import React from 'react'
import Home from './pages/Home'
import Container from './pages/Container'
import Logs from './pages/Logs'

function App() {
  return (
    <Router>
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/containers" element={<Container />} />
            <Route path='/logs/:id' element={<Logs/>} />
          </Routes>
        </main>
    </Router>
  )
}

export default App
