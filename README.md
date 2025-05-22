
# Cloud Data Backup API

## 📘 Intro

This API serves as a cloud data backup system, inspired by Risevest's backend engineer test. It provides authentication and user-based folder management functionalities.

---

## 🔌 Endpoints

### ✅ Authentication

- `POST /auth`  
  **Sign in endpoint**  
  **Status:** ✅ Checked

- `POST /auth/create-account`  
  **Sign up endpoint**  
  **Status:** ✅ Checked

- `GET /me`  
  **Get the current logged in user's data**  
  **Status:** ✅ Checked

---

### 📁 Folder Management

- `GET /folders`  
  **List user's created folders**  
  **Status:** ⛔ Unchecked

- `POST /folder/add`  
  **Create new folder**  
  **Status:** ⛔ Unchecked

- `GET /folder/:id`  
  **Get a specific folder by ID**  
  **Status:** ⛔ Unchecked

- `DELETE /folder/:id`  
  **Delete a specific folder by ID**  
  **Status:** ⛔ Unchecked
