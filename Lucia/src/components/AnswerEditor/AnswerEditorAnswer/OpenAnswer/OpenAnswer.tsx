import TextArea from 'antd/es/input/TextArea';
import { ChangeEvent } from 'react';

interface OpenAnswerProps {
  handleInputChange: (event: ChangeEvent<HTMLTextAreaElement>) => void;
}

export default function OpenAnswer({ handleInputChange }: OpenAnswerProps) {
  return (
    <div>
      <h1>Open Answer:</h1>
      <span>Answer:</span>
      <TextArea name="answer" onChange={handleInputChange} />
      <span>Explanation:</span>
      <TextArea name="explanation" onChange={handleInputChange} />
    </div>
  );
}
