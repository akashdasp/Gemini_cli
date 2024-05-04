# AI Image Description Tool

This tool uses the Google AI language model to generate detailed descriptions for images. It allows you to provide an image and a prompt to describe the image in detail.

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/your-username/your-repo-name.git
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Set up your Google API key:

   - Create a `.env` file in the project directory.
   - Add your Google API key to the `.env` file:

     ```sh
     Google_Api=YOUR_API_KEY_HERE
     ```
4. Another way of setting up your Google API key:

   - pass you api key as the argument -api_key
     ```sh
     go run .\main.go -api_key your_api_key
     ```
   - After this your api key will be saved in a .env file
   - After this you don't have to pass the api key that key will be used 
## Usage

Run the program with the following command:

```sh
go run main.go -image_path=path/to/your/image.png -prompt="Describe the image in detail"
```

Replace `path/to/your/image.png` with the path to your image file and `"Describe the image in detail"` with your desired prompt.
## Usage for Normal Q/A 
Run the program with the following command for Normal Q/A :  

```sh
go run main.go -prompt="What is the meaning of life ?"
```

## Configuration

You can configure the tool using command-line flags:

- `-api_key`: Specify your Google API key.
- `-image_path`: Path to the image file.
- `-prompt`: Prompt to describe the image (default is "Describe the image with detailed manner").

## Contributing

Contributions are welcome! Please follow these steps to contribute:

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
