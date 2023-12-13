import TextArea from 'antd/es/input/TextArea';
import { ChangeEvent } from 'react';

interface SingleChoiceAnswer {
  handleInputChange: (event: ChangeEvent<HTMLTextAreaElement>) => void;
}

export default function SingleAnswer({ handleInputChange }: SingleChoiceAnswer) {
  return (
    <div>
      <h1>Unique answer:</h1>
      <br />
      <TextArea name="answer" onChange={handleInputChange} />
      <span>Explanation:</span>
      <TextArea name="explanation" onChange={handleInputChange} />
    </div>
  );
}
