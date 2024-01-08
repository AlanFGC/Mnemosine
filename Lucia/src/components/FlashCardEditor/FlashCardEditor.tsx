import { useEffect, useState } from 'react';
import Loader from '../UtilityUIcomponents/Loader/Loader';
import FlashCardContextProvider from '../Contexts/FlashCardEditorContext';
import { EditorState } from '../Reducers/FlashcardEditorReducer';
import AnswerEditor from '../AnswerEditor/AnswerEditor';

type FlashCardEditorProps = {
  flashcardID: string | undefined;
};

function FlashCardEditor({ flashcardID = undefined }: FlashCardEditorProps) {
  // eslint-disable-next-line @typescript-eslint/no-unused-vars
  const [fetchedState, setFetchedState] = useState<undefined | EditorState>(undefined);
  const [isLoading, setLoading] = useState(true);
  useEffect(
    () => {
      if (flashcardID) {
        console.log('Loading flashcard');
        // TODO FETCH FROM SERVER
      } else {
        setLoading(false);
      }
    },
    [flashcardID, isLoading, setLoading],
  );
  return (
    isLoading ? <Loader />
      : (
        <FlashCardContextProvider propEditorState={fetchedState}>
          <AnswerEditor prompt="Type your answer" />
        </FlashCardContextProvider>
      )
  );
}

export default FlashCardEditor;
