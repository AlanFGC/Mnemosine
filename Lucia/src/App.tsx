import { Button } from 'antd';
import { useCallback, useState } from 'react';
import CardEditor from './Components/FlashCardEditor/CardEditor';

function App() {
  const [cardEditorValue, setCardEditorValue] = useState('');
  const handleCardEditorChange = useCallback((text: string) => {
    setCardEditorValue(text);
  }, [setCardEditorValue]);

  const handleButtonClick = useCallback(() => {
    console.log(cardEditorValue);
  }, [cardEditorValue]);

  return (
    <div className="App">
      <h1>Welcome to Mnesomine</h1>
      <CardEditor onTextChange={handleCardEditorChange} />
      <Button type="primary" onClick={handleButtonClick}>Button</Button>
    </div>
  );
}

export default App;
