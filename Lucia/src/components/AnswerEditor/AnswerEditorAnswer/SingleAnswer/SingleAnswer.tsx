import Input from 'antd/es/input/Input';

interface SingleChoiceAnswer {
  handleAnswerChange: () => void;
}


export default function SingleAnswer( { handleAnswerChange }: SingleChoiceAnswer) {
  return (
    <div>
      <h1>Unique answer:</h1>
      <br />
      <Input name="answer" />
      <span>Explanation:</span>
      <Input name="explanation" />
    </div>
  );
}
