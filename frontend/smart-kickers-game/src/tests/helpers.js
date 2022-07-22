import { screen } from '@testing-library/react';

export function getElementWhichContain(text) {
  return screen.getByText(text, { exact: false });
}
