import { BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import React from 'react'
import Home from './pages/Home'
import Container from './pages/Container'
import Logs from './pages/Logs'
import ContainerDetails from './pages/ContainerDetails'
import Images from './pages/images/Images'

function App() {
  return (
    <Router>
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/containers" element={<Container />} />
            <Route path="/containers/:id" element={<ContainerDetails />} />
            <Route path='/logs/:id' element={<Logs/>} />
            <Route path="/images" element={<Images />} />
          </Routes>
        </main>
    </Router>
  )
}

export default App
