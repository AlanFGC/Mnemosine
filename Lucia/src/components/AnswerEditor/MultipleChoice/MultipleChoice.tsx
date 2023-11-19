import { ChangeEvent, useState } from 'react';
import { Button } from 'antd';

type IncorrectAnswer = {
  id: number;
  text: string;
};

export default function MultipleChoice(): JSX.Element {
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
      <span>Answers:</span>
      <input type="text" name="answer" />
      <span>Incorrect choices:</span>

      <div>
        <Button onClick={addIncorrectAnswer}>Add</Button>
        <input type="text" name="incorrectAnswer" value={newIncorrectAnswer} onChange={onChangeIncorrectAnswer} />
      </div>

      <ul>

        {incorrectAnswerList.map((answer: IncorrectAnswer) => (
          <>
            <li key={answer.id}>{answer.text}</li>
            <Button onClick={() => removeItem(answer.id)}>X</Button>
          </>
        ))}

      </ul>
      <input type="text" name="answer" />
      <span>Explanation:</span>
      <input type="text" name="explanation" />
    </div>
  );
}
