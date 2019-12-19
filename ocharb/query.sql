SELECT CONCAT_WS(' | ',
        sm.title,
        COALESCE(sm.title_info, ''),
        COALESCE(sm.note, ''),
        COALESCE(sm.comment, ''),
        COALESCE(sm.mark, ''),
        COALESCE(sm.location, '')) text
    FROM sheet_music sm
UNION
SELECT CONCAT_WS(' | ',
        smi.title,
        COALESCE(smi.revision, ''),
        COALESCE(smi.ext_composers, ''),
        COALESCE(smi.ext_authors, ''),
        COALESCE(smi.ext_instruments, '')) text
    FROM sheet_music_items smi;