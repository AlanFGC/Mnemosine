import { Button } from 'antd';
import FlashCardEditor from './components/FlashCardEditor/FlashCardEditor';
import AnswerEditor from './components/AnswerEditor/AnswerEditor';

function App() {
  return (
    <div className="App">
      <h1>Welcome to Mnesomine</h1>
      <FlashCardEditor />
      <AnswerEditor field={0} />
    </div>
  );
}

export default App;
