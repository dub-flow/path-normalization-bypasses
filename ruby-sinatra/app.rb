require 'sinatra'

set :port, 5000
set :bind, '0.0.0.0'

# Root route - Hello World
get '/' do
  'Hello World from Ruby!'
end

# Admin route
get '/admin' do
  'Ruby Admin area'
end
