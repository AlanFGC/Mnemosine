import { ChangeEvent, useState } from 'react';
import { Button } from 'antd';
import TextArea from 'antd/es/input/TextArea';
import { ANSWERS, EXPLANATION, INCORRECTANSWERS } from '../dtypes';

interface MultipleChoiceProps {
  handleInputChange: ((name: string, value: string[]) => void);
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
    handleInputChange(INCORRECTANSWERS, incorrectAnswerList.map((a) => a.text));
  };

  const removeItem = (id: number) => {
    setIncorrectAnswersList(incorrectAnswerList.filter((answer) => answer.id !== id));
    handleInputChange(INCORRECTANSWERS, incorrectAnswerList.map((a) => a.text));
  };

  return (
    <div>
      <h1>Multiple Choice:</h1>
      <br />
      <span>Answer:</span>
      <TextArea
        name={ANSWERS}
        onChange={(e) => handleInputChange(e.target.name, [e.target.value])}
      />
      <span>Incorrect choices:</span>

      <div>
        <Button onClick={addIncorrectAnswer}>Add</Button>
        <input type="text" name={INCORRECTANSWERS} value={newIncorrectAnswer} onChange={(e: ChangeEvent<HTMLInputElement>) => onChangeIncorrectAnswer(e)} />
      </div>

      <ul>

        {incorrectAnswerList.map((answer: IncorrectAnswer) => (
          <li key={answer.id}>
            {answer.text}
            <Button onClick={() => removeItem(answer.id)}>X</Button>
          </li>
        ))}

      </ul>
      <span>Explanation:</span>
      <TextArea rows={4} placeholder="Addiontal information here" name={EXPLANATION} onChange={(e) => handleInputChange(e.target.name, [e.target.value])} />
    </div>
  );
}
