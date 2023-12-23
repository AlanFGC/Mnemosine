import { useEffect, useState } from 'react';

// eslint-disable-next-line @typescript-eslint/ban-types
const Debounce = (value: Function, delay: number) => {
  const [debounceValue, setDebounceValue] = useState(value);
  useEffect(() => {
    const timeout = setTimeout(() => {
      setDebounceValue(value);
    }, delay);
    return () => clearTimeout(timeout);
  }, [delay, value]);

  return debounceValue;
};

export default Debounce;
