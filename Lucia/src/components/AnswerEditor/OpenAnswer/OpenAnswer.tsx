import TextArea from 'antd/es/input/TextArea';

export default function OpenAnswer(): JSX.Element {
  return (
    <div>
      <h1>Open Answer:</h1>
      <span>Answer:</span>
      <TextArea name="answer" />
      <span>Explanation:</span>
      <TextArea name="explanation" />
    </div>
  );
}
