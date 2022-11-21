import '../App/App.css';
import HouseCard from '../Component/HouseComponent/HouseCard/HouseCard';
import NavbarTop from '../Component/NavbarComponent/NavbarTop';

function App() {
  return (
    <div className="App">
      <NavbarTop />
      <div>
        <HouseCardSlide />
      </div>
    </div>
  );
}

export default App;
