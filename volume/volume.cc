#include <alsa/asoundlib.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>

extern "C" {
#include "volume.h"
}

snd_mixer_elem_t *getElement(snd_mixer_t *handle, const char *card, const char *mixer) {
  int err = 0;

  if ((err = snd_mixer_attach(handle, card)) < 0) {
    fprintf(stderr, "Mixer attach %s error: %s", card, snd_strerror(err));
    return NULL;
  }

  if ((err = snd_mixer_selem_register(handle, NULL, NULL)) < 0) {
    fprintf(stderr, "Mixer register error: %s", snd_strerror(err));
    return NULL;
  }

  if ((err = snd_mixer_load(handle)) < 0) {
    fprintf(stderr, "Mixer %s load error: %s", card, snd_strerror(err));
    return NULL;
  }

  snd_mixer_selem_id_t *sid;
  snd_mixer_selem_id_alloca(&sid);
  snd_mixer_selem_id_set_name(sid, mixer);

  snd_mixer_elem_t *element;
  if (!(element = snd_mixer_find_selem(handle, sid))) {
    fprintf(stderr, "Unable to find simple control '%s',%i\n",
            snd_mixer_selem_id_get_name(sid),
            snd_mixer_selem_id_get_index(sid));
    return NULL;
  }

  return element;
}

bool setVolumeForHandle(snd_mixer_t *handle, const char *card, int percentage, const char *mixer) {
  snd_mixer_elem_t *element = getElement(handle, card, mixer);
  if (element == NULL) {
    fprintf(stderr, "Unable to get element.");
    return false;
  }

  for (int i = 0; i <= SND_MIXER_SCHN_LAST; i++) {
    snd_mixer_selem_channel_id_t channel = (snd_mixer_selem_channel_id_t) i;
    if (snd_mixer_selem_has_playback_channel(element, channel)) {
      if (!snd_mixer_selem_has_playback_volume(element)) {
        fprintf(stderr, "Mixer does not have playback volume.");
        return false;
      }

      int err = 0;
      long min, max;
      if ((err = snd_mixer_selem_get_playback_volume_range(element, &min, &max)) < 0) {
        fprintf(stderr, "Unable to get volume range: %s", snd_strerror(err));
        return false;
      }

      long val = ceil(percentage * (max - min) * 0.01 + min);

      if ((err = snd_mixer_selem_set_playback_volume(element, channel, val)) != 0) {
        fprintf(stderr, "Unable to set volume: %s", snd_strerror(err));
        return false;
      }
    }
  }

  return true;
}

bool getVolumeForHandle(snd_mixer_t *handle, const char *card, int *percentage, const char *mixer) {
  snd_mixer_elem_t *element = getElement(handle, card, mixer);
  if (element == NULL) {
    fprintf(stderr, "Unable to get element.");
    return false;
  }

  for (int i = 0; i <= SND_MIXER_SCHN_LAST; i++) {
    snd_mixer_selem_channel_id_t channel = (snd_mixer_selem_channel_id_t) i;
    if (snd_mixer_selem_has_playback_channel(element, channel)) {
      if (!snd_mixer_selem_has_playback_volume(element)) {
        fprintf(stderr, "Mixer does not have playback volume.");
        return false;
      }

      int err = 0;
      long min, max;
      if ((err = snd_mixer_selem_get_playback_volume_range(element, &min, &max)) < 0) {
        fprintf(stderr, "Unable to get volume range: %s", snd_strerror(err));
        return false;
      }

      long val = 0;
      if ((err = snd_mixer_selem_get_playback_volume(element, channel, &val)) != 0) {
        fprintf(stderr, "Unable to get volume: %s", snd_strerror(err));
        return false;
      }

      val = ((val - min) * 100) / (max - min);

      *percentage = (int)val;
      return true;
    }
  }

  fprintf(stderr, "No channels for mixer.");
  return false;
}

int setVolume(char *card, int percentage, char *mixer) {
  snd_mixer_t *handle = NULL;

  int err = 0;
  if ((err = snd_mixer_open(&handle, 0)) < 0) {
    fprintf(stderr, "Mixer open error: %s\n", snd_strerror(err));
    return false;
  }

  bool success = setVolumeForHandle(handle, card, percentage, mixer);

  snd_mixer_close(handle);
  return success ? 1 : 0;
}

int getVolume(char *card, char *mixer) {
  snd_mixer_t *handle = NULL;

  int err = 0;
  if ((err = snd_mixer_open(&handle, 0)) < 0) {
    fprintf(stderr, "Mixer open error: %s\n", snd_strerror(err));
    return false;
  }

  int percentage = -1;
  getVolumeForHandle(handle, card, &percentage, mixer);

  snd_mixer_close(handle);
  return percentage;
}
