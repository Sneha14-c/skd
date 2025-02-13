//import { BrowserRouter,Routes,Route } from 'react-router-dom'
import PageHeader from './header/PageHeader'
import FlightList from './flights/FlightList'
import FlightCreate from './flights/FlightCreate'
import FlightEdit from './flights/FlightEdit'
//import { BrowserRouter, Route, Routes } from 'react-router-dom'
function App() {
  
  return (
    <>
    {/* <BrowserRouter>
      <Routes>
        <Route path="" element={<FlightList/>}/>
        <Route path="/flights/list"={<FlightList/>}/>
        <Route path="/flights/create"={<FlightCreate/>}/>
        <Route path="/flights/edit/:id"={<FlightEdit/>}/>
      </Routes>
    
    </BrowserRouter> */}
       <PageHeader/>
       <FlightList/>
       <FlightCreate/>
       <FlightEdit/>
    

    </>
  )
}

export default App
