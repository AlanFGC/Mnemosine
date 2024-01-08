import TextArea from 'antd/es/input/TextArea';
import { ANSWERS, EXPLANATION } from '../dtypes';

interface SingleChoiceAnswer {
  handleInputChange: ((name: string, value: string[]) => void);
}

export default function SingleAnswer({ handleInputChange }: SingleChoiceAnswer) {
  return (
    <div>
      <h1>Unique answer:</h1>
      <br />
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
