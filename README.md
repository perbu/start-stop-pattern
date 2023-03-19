# Stopping and starting things through a channel

I needed a way to stop and start things through a channel. The thing already took a context, so I made a wrapping 
function that would take a context and a channel and stop and start based on the input from the channel.
