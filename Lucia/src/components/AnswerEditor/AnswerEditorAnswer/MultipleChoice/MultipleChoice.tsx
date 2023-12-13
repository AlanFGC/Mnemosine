import { ChangeEvent, useState } from 'react';
import { Button } from 'antd';
import TextArea from 'antd/es/input/TextArea';

interface MultipleChoiceProps {
  handleInputChange: (event: ChangeEvent<HTMLTextAreaElement>) => void;
}

type IncorrectAnswer = {
  id: number;
  text: string;
};

export default function MultipleChoice({ handleInputChange }: MultipleChoiceProps) {
  const [incorrectAnswerList, setIncorrectAnswersList] = useState<IncorrectAnswer[]>([]);
  const [newIncorrectAnswer, setNewIncorrectAnswer] = useState('');

  const onChangeIncorrectAnswer = (event: ChangeEvent<HTMLInputElement>) => {
    setNewIncorrectAnswer(event.target.value);
  };

  const addIncorrectAnswer = () => {
    if (newIncorrectAnswer.length === 0) {
      return;
    }
    const newAnswer: IncorrectAnswer = { id: Date.now(), text: newIncorrectAnswer };
    setIncorrectAnswersList([...incorrectAnswerList, newAnswer]);
    setNewIncorrectAnswer('');
  };

  const removeItem = (id: number) => {
    setIncorrectAnswersList(incorrectAnswerList.filter((answer) => answer.id !== id));
  };

  return (
    <div>
      <h1>Multiple Choice:</h1>
      <br />
      <span>Answer:</span>
      <TextArea name="answer" onChange={handleInputChange} />
      <span>Incorrect choices:</span>

      <div>
        <Button onClick={addIncorrectAnswer}>Add</Button>
        <TextArea name="incorrectAnswer" value={newIncorrectAnswer} onChange={onChangeIncorrectAnswer} />
      </div>

      <ul>

        {incorrectAnswerList.map((answer: IncorrectAnswer) => (
          <>
            <li key={answer.id}>{answer.text}</li>
            <Button onClick={() => removeItem(answer.id)}>X</Button>
          </>
        ))}

      </ul>
      <span>Explanation:</span>
      <TextArea rows={4} placeholder="Addiontal information here" name="someattribute" onChange={handleAnswerChange} />
    </div>
  );
}
