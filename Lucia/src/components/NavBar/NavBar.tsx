import { AcademicCapIcon } from '@heroicons/react/24/outline';
import { Menu } from 'antd';
import NavBarCSS from './NavbarCSS.module.css';

function Navbar() {
  return (
    <div className={NavBarCSS.bar}>
      <Menu mode="horizontal">
        <Menu.Item key="mail" icon={<AcademicCapIcon className="icon" />}>
          Navigation One
        </Menu.Item>
      </Menu>
    </div>
  );
}

export default Navbar;
