import { Button } from 'antd';
import FlashCardEditor from './components/FlashCardEditor/FlashCardEditor';

function App() {
  return (
    <div className="App">
      <h1>Welcome to Mnesomine</h1>
      <FlashCardEditor />
      <Button type="primary">Button</Button>
    </div>
  );
}

export default App;
