dir="$(pwd)"

export GOOGLE_APPLICATION_CREDENTIALS="${dir}/Secrets/firebase-service-account.json"
export FIREBASE_STORAGE_API="https://firebasestorage.googleapis.com/v0"


source "${dir}/Secrets/credentials.sh"
