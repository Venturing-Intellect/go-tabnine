import React, { useState } from 'react';
import axios from 'axios';
import 'bootstrap/dist/css/bootstrap.min.css';

const FeedbackForm = () => {
  const [email, setEmail] = useState('');
  const [message, setMessage] = useState('');
  const [status, setStatus] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();
    setStatus('Submitting...');

    try {
      const response = await axios.post('/feedback', { email, message });
      setStatus('Feedback submitted successfully!');
      setEmail('');
      setMessage('');
    } catch (error) {
      setStatus(`Error: ${error.response?.data?.error || 'Something went wrong'}`);
    }
  };

  return (
    <div className="container mt-5">
      <h2 className="mb-4">Submit Feedback</h2>
      <form onSubmit={handleSubmit}>
        <div className="mb-3">
          <label htmlFor="email" className="form-label">Email address</label>
          <input
            type="email"
            className="form-control"
            id="email"
            value={email}
            onChange={(e) => setEmail(e.target.value)}
            required
          />
        </div>
        <div className="mb-3">
          <label htmlFor="message" className="form-label">Feedback</label>
          <textarea
            className="form-control"
            id="message"
            rows="3"
            value={message}
            onChange={(e) => setMessage(e.target.value)}
            required
          ></textarea>
        </div>
        <button type="submit" className="btn btn-primary">Submit</button>
      </form>
      {status && <div className="mt-3 alert alert-info">{status}</div>}
    </div>
  );
};

export default FeedbackForm;