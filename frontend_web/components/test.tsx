import React from 'react';
import fs from 'fs'

const Test = ({ data }: { data: string }) => {
  // Write data to file
  fs.writeFileSync('content/examPrep.md', JSON.stringify(data, null, 2));
  return (
    <div>
    
    </div>
  );
};

export default Test;
