import Input from 'antd/es/input/Input';

export default function SingleAnswer(): JSX.Element {
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
