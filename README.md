# 2fa

Just a simple example code in Go showing a minimal end-to-end flow for server side support for google authenticator. The example code sets up a secret for a single user, generates a QR code as a PNG file that the user can scan into Google Authenticator, and then prompts the user for a token that the user copies from the Authenticator app. We validate the token and print out whether it is valid or not.
