import TextArea from 'antd/es/input/TextArea';

interface OpenAnswerProps {
  handleAnswerChange: () => void;
}

export default function OpenAnswer( { handleAnswerChange }: OpenAnswerProps) {
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
