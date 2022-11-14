import React from 'react'
import 'bootstrap/dist/css/bootstrap.min.css';
import './NavbarTop.css'
import {
  Navbar,
  NavLink,
  Dropdown,
  DropdownToggle,
  DropdownMenu,
  DropdownItem
} from 'reactstrap';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'
import { faUser } from '@fortawesome/free-solid-svg-icons'

const element = <FontAwesomeIcon icon={faUser} />

function NavbarTop() {
  // Collpase isOpen State
  const [isToggle, setIsToggle] = React.useState(false);
  return (
    <div>
      <Navbar color="dark">
        <div className="col-sm-1">
          <Dropdown isOpen={isToggle} toggle={() => { setIsToggle(!isToggle) }}>
            <DropdownToggle caret>
              {element}
            </DropdownToggle>
            <DropdownMenu>
              <DropdownItem>Another Action</DropdownItem>
              <DropdownItem divider />
              <DropdownItem>Profile</DropdownItem>
            </DropdownMenu>
          </Dropdown>
        </div>
        <div className="col-sm-3">
          <ul className="NavList">
            <li>
              <NavLink href="#">Manager</NavLink>
            </li>
            <li>
              <NavLink href="#">Chatting</NavLink>
            </li>
            <li>
              <NavLink href="#">Data</NavLink>
            </li>
          </ul>
        </div>
        <div className="col-sm-4">
          <div className="NavTitle">
            <h5 href="/">Your House</h5>
          </div>
        </div>
        <div className="col-sm-4 from-group">
          <form>
            <div className="row">
              <div className="col-sm-8">
                <input className="form-control" placeholder="Search something you need..."></input>
              </div>
              <div className="col-sm-4">
                <button className="btn btn-primary" type="button">Seach</button>
              </div>
            </div>
          </form>
        </div>
      </Navbar >
    </div >
  );
}

export default NavbarTop
