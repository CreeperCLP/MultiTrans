import React from 'react';

const App: React.FC = () => {
    return (
        <div style={{ 
            display: 'flex', 
            flexDirection: 'column', 
            alignItems: 'center', 
            justifyContent: 'center', 
            height: '100vh', 
            backgroundColor: 'rgba(255, 255, 255, 0.8)', 
            backdropFilter: 'blur(10px)', 
            borderRadius: '10px', 
            boxShadow: '0 4px 30px rgba(0, 0, 0, 0.1)' 
        }}>
            <h1>Clipboard Transfer Station</h1>
            <p>Welcome to the Clipboard Transfer Station!</p>
            <p>Your clipboard data is securely transferred across devices.</p>
        </div>
    );
};

export default App;