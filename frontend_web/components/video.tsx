import React from 'react';

const VideoComponent = () => {
  return (<div className="relative overflow-hidden w-full" style={{ paddingTop: '56.25%' }}>
      <video
        className="absolute inset-0 w-full h-full object-cover"
        src="https://d3c33hcgiwev3.cloudfront.net/VP8oxDdSSIe_KMQ3UmiHfg.processed/full/240p/index.mp4?Expires=1713139200&Signature=UfTWNa9qCw9siq3~tevWhepZ5MQQHZy6LeTvOE4ob5KzHGFvIZlZNk-P4OQKnDrCyNPj-wxKHdFaBmSOIxqBKGFwQgzvhSrciaiw6qQrhtj~fHbkzhGQeXHv-V33xMQ-qR-pmp4H1WkwpNT8188OLpdBDJ5UGcEDNLeonW2AxOg_&Key-Pair-Id=APKAJLTNE6QMUY6HBC5A"
        controls
      ></video>
    </div>
  );
};

export default VideoComponent;
