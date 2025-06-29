import { BrowserRouter as Router, Routes, Route} from 'react-router-dom'
import React from 'react'
import Home from './pages/Home'
import Container from './pages/Container'

function App() {
  return (
    <Router>
        <main>
          <Routes>
            <Route path="/" element={<Home />} />
            <Route path="/containers" element={<Container />} />
          </Routes>
        </main>
    </Router>
  )
}

export default App
