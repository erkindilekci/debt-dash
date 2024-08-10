# Debt Dash
Debt Dash is a web application designed for individuals residing in Turkey who manage multiple credit cards. The app offers a comprehensive platform to monitor and manage outstanding credit card balances, including both current debt amounts and associated minimum payments.

## Technologies

- `Go`
- `Chi Router`
- `PostgreSQL`
- `Supabase`

## Setting Up Debt Dash Locally

To run the Debt Dash application on your local machine, follow these steps:

1. **Clone the Repository:**
   Begin by cloning the Debt Dash repository to your local environment using Git:
   ```bash
   https://github.com/erkindilekci/debt-dash.git
   ```

2. **Configure Database Connection:**
   Edit the `db.go` file located in the `internal/db` directory. Replace the placeholder PostgreSQL URL with your specific database connection string.

3. **Build the Docker Image:**
   Create a Docker image from the project's Dockerfile:
   ```bash
   docker build -t debt-dash:1.0 .
   ```

4. **Run the Docker Container:**
   Start the Docker container, mapping port 8080 of the container to port 8080 on your host:
   ```bash
   docker run -d -p 8080:8080 debt-dash:1.0
   ```

5. **Access the Application:**
   Open your web browser and navigate to `http://localhost:8080` to access the Debt Dash application.

**Note:** Ensure you have Docker Desktop installed and running on your system before proceeding with these steps.


## Screenshoots
![debtdash](https://github.com/user-attachments/assets/de15fc29-d7d1-4a09-82b4-f27f8a05ec0e)
