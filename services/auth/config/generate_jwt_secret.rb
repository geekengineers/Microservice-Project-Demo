require "securerandom"

secret = SecureRandom.urlsafe_base64 100

File.open(File.join(File.dirname(__FILE__), "jwt_secret.pem"), "w") do |f|
  f.write secret
  f.close
end