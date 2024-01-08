import React, { useState } from 'react';
import axios from 'axios';

const InstanceDetails = () => {
  const [instanceId, setInstanceId] = useState('');
  const [responseData, setResponseData] = useState(null);

  const handleGetInstanceDetails = async () => {
    try {
      const response = await axios.get(`http://localhost:8080/describe-instance?id=${instanceId}`);
      setResponseData(response.data);
    } catch (error) {
      console.error('Error fetching instance details:', error);
    }
  };

  return (
    <div>
      <h1>EC2 Instance Details</h1>
      <label>
        Enter Instance ID:
        <input type="text" value={instanceId} onChange={(e) => setInstanceId(e.target.value)} />
      </label>
      <button onClick={handleGetInstanceDetails}>Get Instance Details</button>

      {responseData && (
        <div>
          <h2>Instance Details:</h2>
          <pre>{JSON.stringify(responseData, null, 2)}</pre>
        </div>
      )}
    </div>
  );
};

export default InstanceDetails;
