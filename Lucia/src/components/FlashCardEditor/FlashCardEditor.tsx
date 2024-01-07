import { useEffect, useState } from 'react';
import Loader from '../UtilityUIcomponents/Loader/Loader';
import FlashCardContextProvider from '../Contexts/FlashCardEditorContext';
import { EditorState } from '../Reducers/FlashcardEditorReducer';

type FlashCardEditorProps = {
  flashcardID: string | undefined;
};

function FlashCardEditor({ flashcardID = undefined }: FlashCardEditorProps) {
  const [fetchedState, setFetchedState] = useState<undefined | EditorState>(undefined);
  const [isLoading, setLoading] = useState(true);
  useEffect(
    () => {
      if (flashcardID) {
        console.log('Loading flashcard');
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
          <h1>Test</h1>
        </FlashCardContextProvider>
      )
  );
}

export default FlashCardEditor;
