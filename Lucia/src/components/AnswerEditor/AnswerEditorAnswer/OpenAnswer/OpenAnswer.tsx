import TextArea from 'antd/es/input/TextArea';
import { ANSWERS, EXPLANATION } from '../dtypes';

interface OpenAnswerProps {
  handleInputChange: ((name: string, value: string[]) => void);
}

export default function OpenAnswer({ handleInputChange }: OpenAnswerProps) {
  return (
    <div>
      <h1>Open Answer:</h1>
      <span>Answer:</span>
      <TextArea
        name={ANSWERS}
        onChange={(e) => handleInputChange(e.target.name, [e.target.value])}
      />
      <span>Explanation:</span>
      <TextArea
        name={EXPLANATION}
        onChange={(e) => handleInputChange(e.target.name, [e.target.value])}
      />
    </div>
  );
}
